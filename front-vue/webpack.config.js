const path = require('path')
const webpack = require('webpack')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const CleanWebpackPlugin = require('clean-webpack-plugin')
const ExtractTextPlugin = require('extract-text-webpack-plugin')
const StatsPlugin = require('stats-webpack-plugin')
const SriPlugin = require('webpack-subresource-integrity')

module.exports = function (options) {
  var entry = {
    app: path.join(__dirname, 'app')
  }
  entry.vendor = [
    // 'bootstrap',
    'tether',
    'bootstrap',
    /* redux */
    'redux',
    'revue',
    /* vue */
    'vue',
    'vue-resource',
    'vue-router',
    'pug',
    /* utils*/
    'jwt-decode'
  // 'url-parse',
  // 'marked',
  // i18n
  // 'i18next',
  // 'i18next-xhr-backend',
  // 'i18next-browser-languagedetector'
  ]

  var plugins = [
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      'window.Tether': 'tether'
    })]
  var loaders = [{
    test: /\.vue$/,
    loader: 'vue'
  }, {
    test: /\.js$/,
    exclude: /(node_modules)/,
    loader: 'babel'
  }, {
    test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/,
    loader: 'file'
  }, {
    test: /\.json$/,
    loader: 'json'
  }]

  var env = {
    CHAOS: JSON.stringify({
      backend: options.backend,
      engines: options.engines,
      version: require('./package.json').version
    })
  }
  var output = {
    path: path.join(__dirname, 'dist'),
    publicPath: '/'
  }
  var htmlOptions = {
    inject: true,
    template: 'app/index.html',
    filename: 'index.html',
    favicon: path.join(__dirname, 'app', 'favicon.png'),
    title: 'IT-PACKAGE'
  }

  if (options.minify) {
    env['process.env.NODE_ENV'] = JSON.stringify('production')
    output.filename = '[id]-[chunkhash].js'
    htmlOptions.minify = {
      collapseWhitespace: true,
      removeComments: true
    }

    plugins.push(new CleanWebpackPlugin(['build']))
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
    }))
    plugins.push(new webpack.optimize.CommonsChunkPlugin({
      name: 'vendor',
      minChunks: 3
    }))
    plugins.push(new webpack.optimize.DedupePlugin())
    plugins.push(new webpack.optimize.OccurrenceOrderPlugin(true))
    plugins.push(new webpack.NoErrorsPlugin())
    plugins.push(new ExtractTextPlugin('[id]-[chunkhash].css'))
    plugins.push(new SriPlugin(['sha256', 'sha512']))

    loaders.push({
      test: /\.css$/,
      loader: ExtractTextPlugin.extract('style', 'css')
    })
  } else {
    output.filename = '[name].js'

    plugins.push(new webpack.SourceMapDevToolPlugin({}))
    plugins.push(new StatsPlugin('stats.json', {
      chunkModules: true,
      exclude: [/node_modules/]
    }))
    loaders.push({
      test: /\.css$/,
      loaders: ['style', 'css']
    })
  }

  plugins.push(new webpack.DefinePlugin(env))
  plugins.push(new HtmlWebpackPlugin(htmlOptions))

  return {
    entry: entry,
    output: output,
    // resolveLoader: {
    //   root: path.join(__dirname, 'node_modules')
    // },
    plugins: plugins,
    module: {
      // preLoaders: [{
      //   test: /\.js$/,
      //   loader: 'eslint-loader',
      //   exclude: /node_modules/
      // }],
      loaders: loaders
    },
    node: {
      // fix: Cannot find module "fs"
      fs: 'empty'
    },
    devServer: {
      historyApiFallback: true,
      port: 4200
    }
  }
}
