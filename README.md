
Request-Promise(BlurBird) [![build status](http://img.shields.io/travis/reactjs/react-redux/master.svg?style=flat-square)](http://www.guozj.com)
=========================

This is an easy example about how to manage your restful end points efficiently, with a uniform requestor working with blurbird promise.


## cli example

```
npm install
```
Install dependency
```
node server.js
```
Run express example at port 8080
```
npm run watch
```


## Highlight

- Request
- BlurBird

## How Does It Work?


For instance, one end point is [https://registry.npmjs.org/react](https://registry.npmjs.org/react).

It is recommended to add the base restful Url in the environment variable, in this case, `https://registry.npmjs.org`

The api method here is GET, and looks like `/{moduleName} ` , when replace moduleName with `react `

Follow the example in the controller so that you could easily create the bluebird promise and manage the endpoint very efficiently with less trouble in naming your endpoint.

## Customize and global control your request

You are free to make a global setting on the request class, ref to [options](https://github.com/request/request)

Meanwhile, you could also make any local override request configuration on each individual request instance.

Take `timeout` for example, the default in the repo is set to 20000, in the individual variable, you could change to whatever new value based on this need.


## License

MIT