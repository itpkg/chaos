import React from 'react'

const Widget = React.createClass({
  render() {
    return (
      <div>
        <h1>App</h1>
        {this.props.children}
      </div>
    )
  }
})

export default Widget
