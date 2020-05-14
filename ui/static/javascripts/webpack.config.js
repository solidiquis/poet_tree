module.exports = {
  mode: "production",
  devtool: "source-map",
  resolve: {
      extensions: [".ts", ".tsx", ".js", ".css"]
  },
  module: {
      rules: [
          {
              test: /\.ts(x?)$/,
              exclude: /node_modules/,
              use: [
                  {
                      loader: "ts-loader"
                  }
              ]
          },
          {
              enforce: "pre",
              test: /\.js$/,
              loader: "source-map-loader"
          },
          { 
              test: /\.css$/,
              use: [
                  "style-loader",
                  "css-loader"
              ]
          },
      ]
  },
  externals: {
      "react": "React",
      "react-dom": "ReactDOM"
  }
};