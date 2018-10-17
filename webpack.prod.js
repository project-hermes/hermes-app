const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
    mode: 'production',
    plugins: [new UglifyJsPlugin()],
    resolve: {
        alias: {
            vue: 'vue/dist/vue.min'
        }
    }
});
