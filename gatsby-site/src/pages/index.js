import React, { Component } from "react"
import { Link, graphql } from "gatsby"

import Layout from "../components/layout"


export default class Index extends Component {
  constructor(props) {
    super()
    console.log(props)
    this.state = {
      location: props.location,
      data: props.data,
      siteTitle: props.data.site.siteMetadata.title
    }
  }
  render(){
    return (<Layout class="index" location={this.state.location} title={this.state.siteTitle} >Hello
    
      <Link className="header-link-home" to="/home">
        HomePage
      </Link>
      </Layout>)
  }
}

export const pageQuery = graphql`
  query {
    site {
      siteMetadata {
        title
      }
    }
    allMarkdownRemark(sort: { fields: [frontmatter___date], order: DESC }) {
      nodes {
        excerpt
        fields {
          slug
        }
        frontmatter {
          date(formatString: "MMMM DD, YYYY")
          title
          description
        }
      }
    }
  }
`
