const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const StatsPlugin = require("stats-webpack-plugin");

const VERSION = '2016.6.9';

module.exports = function(options) {
    var entry = options.engines.reduce(function(obj, en) {
        obj[en] = path.join(__dirname, "app", "engines", en);
        return obj
    }, {});
    entry.vendor = [
        'jquery',
        'bootstrap',

        'react',
        'react-dom',
        'react-router',
        'react-bootstrap',
        'react-redux',
        'react-router-redux',

        'jwt-decode',
        'url-parse',
        'marked'
    ];

    var plugins = [
        new HtmlWebpackPlugin({
            title: 'Webpack demo'
        })
    ];
    var loaders = [{
        test: /\.jsx?$/,
        exclude: /(node_modules)/,
        loader: "babel"
    }, {
        test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/,
        loader: "file"
    }];

    var env = {
        'app.host': JSON.stringify(options.host),
        'app.engines': JSON.stringify(options.engines),
        'app.version': JSON.stringify(VERSION)
    };
    var output = {
        path: path.join(__dirname, 'build')
    };
    var htmlOptions = {
        inject: true,
        template: 'app/index.html',
        favicon: path.join(__dirname, 'app', 'favicon.png')
    };

    if (options.minify) {
        env['process.env.NODE_ENV'] = 'production';
        output.filename = '[id]-[chunkhash].js';
        htmlOptions.minify = {
            collapseWhitespace: true,
            removeComments: true
        };

        plugins.push(new CleanWebpackPlugin(['build']));
        plugins.push(new webpack.optimize.UglifyJsPlugin({
            compress: {
                drop_console: true,
                drop_debugger: true,
                // dead_code: true,
                // unused: true,

                warnings: false
            },
            output: {
                comments: false
            }
        }));
        plugins.push(new webpack.optimize.CommonsChunkPlugin({
            name: "vendor",
            minChunks: 3
        }));
        plugins.push(new webpack.optimize.DedupePlugin());
        plugins.push(new webpack.optimize.OccurrenceOrderPlugin(true));
        plugins.push(new webpack.NoErrorsPlugin());
        plugins.push(new ExtractTextPlugin('[id]-[chunkhash].css'));

        loaders.push({
            test: /\.css$/,
            loader: ExtractTextPlugin.extract('style', 'css')
        });
    } else {
        output.filename = '[name].js';

        plugins.push(new webpack.SourceMapDevToolPlugin({}));
        plugins.push(new StatsPlugin('stats.json', {
            chunkModules: true,
            exclude: [/node_modules/]
        }));
        loaders.push({
            test: /\.css$/,
            loaders: ['style', 'css']
        });
    }

    plugins.push(new webpack.DefinePlugin(env));
    options.engines.forEach(function(en) {
        plugins.push(new HtmlWebpackPlugin(Object.assign({},
            htmlOptions, {
                title: en,
                filename: (en == 'platform' ? 'index' : en) + ".html",
                chunks: ['vendor', en]
            }
        )));
    });

    return {
        entry: entry,
        output: output,
        plugins: plugins,
        module: {
            loaders: loaders,
        },
        resolve: {
            extensions: ['', '.js', '.jsx'],
        },
        devServer: {
            //historyApiFallback: true,
            port: 4200
        }
    };
}