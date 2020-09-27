import React from "react"
import { Link } from "gatsby"

import Bio from "../components/bio"
import Layout from "../components/layout"
import SEO from "../components/seo"

const BlogIndex = (props) => {
    const siteTitle = props.data.site.siteMetadata?.title || `Title`
    const posts = props.data.allMarkdownRemark.nodes
    const location = props.location
    console.log(props)

    return (
        <Layout location={location} title={siteTitle}>
            <SEO title="All posts" />
            <Bio />
            {posts.length === 0
                ?
                <Layout location={location} title={siteTitle}>
                    <SEO title="All posts" />
                    <Bio />
                    <p>
                        No blog posts found. Add markdown posts to "content/blog" (or the
                        directory you specified for the "gatsby-source-filesystem" plugin in
                        gatsby-config.js).
                    </p>
                </Layout>
                :
                posts.map(post => {
                    const title = post.frontmatter.title || post.fields.slug
                    return (
                        <article
                            key={post.fields.slug}
                            className="post-list-item"
                            itemScope
                            itemType="http://schema.org/Article"
                        >
                            <header>
                                <h2>
                                    <Link to={post.fields.slug} itemProp="url">
                                        <span itemProp="headline">{title}</span>
                                    </Link>
                                </h2>
                                <small>{post.frontmatter.date}</small>
                            </header>
                            <section>
                                <p
                                    dangerouslySetInnerHTML={{
                                        __html: post.frontmatter.description || post.excerpt,
                                    }}
                                    itemProp="description"
                                />
                            </section>
                        </article>
                    )
                })}
        </Layout>
    )
}

export default BlogIndex

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