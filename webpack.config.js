/* eslint no-use-before-define: 0 */

let path = require('path');

let conf = {
    entry: "./js_modules/index.js",
    output: {
        path: path.resolve(__dirname, "./"),
        filename: "main.js",
        publicPath: "js/"
    },
    devServer: {
        overlay: true
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                loader: "babel-loader",
                exclude: /node_modules/
            }
        ]
    }
};

module.exports = (env, options) => {
    conf.devtool = options.mode === "production" 
                                ? false 
                                : "cheap-module-eval-source-map";
    return conf;
};