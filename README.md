# go_vue_template

## Requirements

* Install [Go (Programming Language)](https://go.dev/dl/)
* Install [Node Package Manager (NPM)](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
* Install [OpenSSL](https://www.openssl.org/source/)
  * Quick installation of openssl under Debian: `sudo apt install openssl`
* Add SSL certificate and key to root of project as `cert.pem` and `key.pem`
  * You can self sign your own key/cert pair like this: 
      ```
        openssl req \
        -newkey rsa:4096 -nodes -keyout key.pem \
        -x509 -days 365 -out cert.pem```
  * If you don't want to use SSL, look in `cmd/main.go` and swap http listener.
* Build the UI with npm
  * Navigate into the `ui` folder
  * execute `npm i`.
    * This installs any dependencies of the frontend UI.
  * execute `npm run build`.
    * Must be done each time you make changes to the frontend. (hot-reload command in ui/readme.md)
    * This builds the frontend code into a distribution folder the backend server can host.
  * navigate back to the parent repository folder.
* Create the .env file
  * example: 
        ```
          JWT_SECRET="SUPERSECRET!"
        ```
