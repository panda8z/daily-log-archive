import React, { Component } from 'react';

export default class PostFooter extends Component {
    constructor(props) {
        super()
    }

    /**
     *  作者： Changkun Ou
链接： https://blog.changkun.de/archives/2019/06/261/
版权： All posts in this blog are licensed under CC BY-NC-ND 4.0.
     */
    render() {
        return (<div class="poster-footer-lisence">
            <a href="https://creativecommons.org/licenses/by-nc-nd/4.0/">
                <img src="https://mirrors.creativecommons.org/presskit/buttons/88x31/png/by-nc-nd.png"/>
                footer
            </a>
        </div>)
    }
}
