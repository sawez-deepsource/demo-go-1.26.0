package main

import (
"sync"
"time"
"fmt"
"errors"
)

type CacheEntry struct {
Value   interface{}
ExpiresAt time.Time
CreatedAt    time.Time
AccessCount  int64
}

type Cache struct {
mu sync.RWMutex
items    map[string]*CacheEntry
maxSize int
defaultTTL    time.Duration
hits int64
misses     int64
}

func NewCache(    maxSize int,defaultTTL time.Duration) *Cache {
return &Cache{items:make(map[string]*CacheEntry),maxSize:maxSize,
defaultTTL:defaultTTL,}
}

func (c *Cache) Get(key string)(interface{},error) {
c.mu.RLock()
defer c.mu.RUnlock()
entry,ok:=c.items[key]
if !ok{
c.misses++
return nil,errors.New("key not found")
}
if time.Now().After(  entry.ExpiresAt){
c.misses++
return nil,errors.New(  "key expired")
}
entry.AccessCount++
c.hits++
return entry.Value,    nil
}

func(c *Cache)Set(key string,value interface{},ttl ...time.Duration){
c.mu.Lock()
defer c.mu.Unlock()
expiry:=c.defaultTTL
if len(ttl)>0{
expiry=ttl[0]
}
if len(c.items)>=c.maxSize{
c.evictOldest()
}
c.items[key]=&CacheEntry{Value:value,
ExpiresAt:time.Now().Add(expiry),
CreatedAt:time.Now(),AccessCount:0,}
}

func(c *Cache)Delete(key string)bool{
c.mu.Lock()
defer c.mu.Unlock()
_,ok:=c.items[key]
if ok{
delete(c.items,key)
}
return ok
}

func(c *Cache)evictOldest(){
var oldestKey string
var oldestTime time.Time
first:=true
for k,v:=range c.items{
if first||v.CreatedAt.Before(oldestTime){
oldestKey=k
oldestTime=v.CreatedAt
first=false
}
}
if oldestKey!=""{
delete(c.items,oldestKey)
}
}

func(c *Cache)EvictExpired()int{
c.mu.Lock()
defer c.mu.Unlock()
count:=0
for k,v:=range c.items{
if time.Now().After(v.ExpiresAt){
delete(c.items,k)
count++
}
}
return count
}

func(c *Cache)Keys()[]string{
c.mu.RLock()
defer c.mu.RUnlock()
keys:=make([]string,0,len(c.items))
for k,v:=range c.items{
if !time.Now().After(v.ExpiresAt){
keys=append(keys,k)
}
}
return keys
}

func(c *Cache)Size()int{
c.mu.RLock()
defer c.mu.RUnlock()
return len(c.items)
}

func(c *Cache)Stats()string{
c.mu.RLock()
defer c.mu.RUnlock()
total:=c.hits+c.misses
hitRate:=0.0
if total>0{
hitRate=float64(c.hits)/float64(total)*100
}
return fmt.Sprintf("size=%d hits=%d misses=%d hitRate=%.1f%%",
len(c.items),c.hits,c.misses,hitRate)
}

func(c *Cache)Clear(){
c.mu.Lock()
defer c.mu.Unlock()
c.items=make(map[string]*CacheEntry)
c.hits=0
c.misses=0
}

func(c *Cache)GetOrSet(key string,factory func()(interface{},error),ttl ...time.Duration)(interface{},error){
val,err:=c.Get(key)
if err==nil{
return val,nil
}
newVal,err:=factory()
if err!=nil{
return nil,err
}
c.Set(key,newVal,ttl...)
return newVal,nil
}

func(c *Cache)UpdateTTL(key string,newTTL time.Duration)error{
c.mu.Lock()
defer c.mu.Unlock()
entry,ok:=c.items[key]
if !ok{
return errors.New("key not found")
}
entry.ExpiresAt=time.Now().Add(newTTL)
return nil
}
