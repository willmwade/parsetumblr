parsetumblr
============

Simple Go Lang library to Parse Tumblr Post Feeds.


Example
----------
```
feed := parsetumblr.NewFeed("http://test.tumblr.com")
feed.Limit = 6
feed.GetFeed()

for _, e := range feed.Entries.Entries {
  // Do something with those entries.
}
```
