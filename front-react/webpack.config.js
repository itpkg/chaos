const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const StatsPlugin = require("stats-webpack-plugin");

module.exports = function(options) {

    var entry = {
        app: path.join(__dirname, "app")
    };
    entry.vendor = [
        'jquery',
        'bootstrap',

        'react',
        'react-dom',
        'react-router',
        'react-bootstrap',
        'react-redux',
        'react-router-redux',
        'react-router-bootstrap',

        'jwt-decode',
        'url-parse',
        'marked',

        'i18next',
        'i18next-xhr-backend',
        'i18next-browser-languagedetector'
    ];

    var plugins = [
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
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
      CHAOS_ENV: JSON.stringify({
        backend:options.backend,
        engines: options.engines,
        version: '2016.6.9'
      })
    };
    var output = {
        path: path.join(__dirname, 'build'),
        publicPath: '/'
    };
    var htmlOptions = {
        inject: true,
        template: 'app/index.ejs',
        filename: 'index.html',
        favicon: path.join(__dirname, 'app', 'favicon.png'),
        title: 'IT-PACKAGE'
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
    plugins.push(new HtmlWebpackPlugin(htmlOptions));

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
            historyApiFallback: true,
            port: 4200
        }
    };
}
