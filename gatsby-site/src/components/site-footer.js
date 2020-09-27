import React, { Component } from 'react';

export default class SiteFooter extends Component {
    constructor(props) {
        super()
    }

    render() {
        return (<footer>
            Corpyright © 2020 Panda张向北. Allrights reserved. <br></br>
            Powered by
            {` `}
            <a href="https://www.gatsbyjs.com">Gatsby</a>
        </footer>)
    }
}
