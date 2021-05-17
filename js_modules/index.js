// eslint-disable-next-line no-unused-vars
import Articles from "./articles-react/articles"
import ReactDOM from "react-dom"
// eslint-disable-next-line no-unused-vars
import React from "react"
import * as Article from "./articles-react/article-model"
import "@babel/polyfill"

async function App()
{
    let articles = await Article.all();
    ReactDOM.render( <Articles articles={articles}/>, document.querySelector('.app'));
}

App();
