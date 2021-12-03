# GCurtain

Easy way to control your features using [redis](http://redis.io/).

Also available for Java -> [JCurtain!](https://github.com/moip/jcurtain)

## Installation

Add this line to your application's Gemfile:

```Go
rcurtain
```

## Usage

* Gcurtain uses redis to control features, which can be checked by a **percentage** or a **set of users**.
```
feature:[name-of-feature]:percentage
```
```
feature:[name-of-feature]:users
```

* To use Gcurtain, first your need to initialize the configuration defining your **redis URL** (password@ip:port/database). Optionally, you can also configure the **default response** when the feature is not found, which by default is false.

```go
func getClient(uri string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: uri,
	})
}
```

* Get the instance of Gcurtain.
```go
const uri = "localhost:6379"

var g = new(GCurtain))
g.Init(uri)
```

* Consult if the curtain is opened for a feature using the method "IsOpen", passing the name of the feature you want to check.
```go
g.IsOpen('feature')
```

* You can also pass a set of users to be checked.
```go
g.IsOpen('feature', ['user-1','user-2'])
```

## How to start unit testing
1. Upload redis image
2. run the gcurtain_test file (go test gcurtain_test)


## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/wirecardBrasil/gcurtain. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

1. Fork it ( https://github.com/wirecardBrasil/gcurtain)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
