import React, { Component } from "react"
import { Link, graphql } from "gatsby"


export default class Index extends Component {
  constructor(props) {
    super()
    console.log(props)
    this.state = {
      location: props.location,
      img: props.data.img.childImageSharp.fixed,
    }
  }
  render() {
    return (<div className="index" >
    
      <div 
      style={{ 
        backgroundImage : this.state.img.src,
        backgroundRepeat: 'repeat',
        display:"flex",
        position: "absolute",
        flexDirection: "row",
      }}>
      <Link className="header-link-home" to="/home">
        Home
      </Link>
      <br/>
      <Link className="header-link-home" to="/list">
        Blog List
      </Link>
      </div>
    </div>)
  }
}

export const pageQuery = graphql`
{
  img: file(absolutePath: {regex: "/profile-pic.jpg/"}) {
    childImageSharp {
      fixed(width: 500, height: 500, quality: 95) {
        ...GatsbyImageSharpFixed
      }
    }
  }
}
`