import React from "react";
import Article from "./article-component";
import * as ArticleModel from "./article-model"
import {Typography, Button, TextField} from "@material-ui/core";
import ModalAddArticle from "./modal-add-article-component";
import ModalViewArticle from "./modal-view-article-component";
import ArticlesAll from "./articles-component";
import ModalEditArticle from "./modal-edit-article-component";
import "@babel/polyfill";
import { server } from "../requests";

class Articles extends React.Component
{
    constructor(obj){
        super(obj);

        let popupInfo = {
            isOpen: false,
            content: {
                title: "",
                body: "",
            },
            article_id: null
        };

        this.state = {
            articles: this.props.articles,
            popups: {}
        };

        let popups = ['addPopup', 'viewOnePopup', 'editOnePopup'];
        for( let popup of popups ){
            this.state.popups[popup] = {...popupInfo}
        }
    }

    async onAdd(data)
    {
        let result = await ArticleModel.add( data );
        let state = {...this.state};
        let articles =  state.articles;

        articles.unshift( result );
        state.popups.addPopup.isOpen = false;
        this.setState(state);
    }

    async onEdit()
    {
        let id = this.state.popups.editOnePopup.article_id;
        let data = this.state.popups.editOnePopup.content;

        let result = await ArticleModel.edit( id+1, data );

        if( result != null ){
            let state = {...this.state};
            Object.assign( state.articles[id], result );
            state.popups.editOnePopup.isOpen = false;
            this.setState(state);
        }
    }

    changePopupState( popupName, bool )
    {
        if( popupName in this.state.popups ){
            let popups = {...this.state.popups};
            popups[popupName].isOpen = bool;
            this.setState( { popups } );
        }
    }

    changeAddPopupState(bool)
    {
        this.changePopupState('addPopup', bool);
    }

    changeViewOnePopup(bool)
    {
        this.changePopupState('viewOnePopup', bool);
    }

    changeEditPopupState(bool)
    {
        this.changePopupState('editOnePopup', bool);
    }

    async onDelete( id )
    {
        if( id == undefined ){ return; }
        let articles =  [...this.state.articles];

        let result = await ArticleModel.remove( id );

        for(var i = 0; i < articles.length; i++){
            if( articles[i].id == id ){
                articles.splice(i, 1);
            }
        }

        this.setState({articles});
    }

    getViewOnePopup()
    {
        let data = this.state.popups.viewOnePopup.content;
        
        return <ModalViewArticle 
            open={this.state.popups.viewOnePopup.isOpen}
            onClose={ () => { this.changeViewOnePopup(false) } } 
            onOk={ () => { this.changeViewOnePopup(false) } } 
            title={ data.title }
            body={ data.body }
            />
    }

    getAddOnePopup()
    {
        return <ModalAddArticle 
            open={this.state.popups.addPopup.isOpen}
            onClose={ () => { this.changeAddPopupState(false) } } 
            onOk = { (article) => { this.onAdd(article) } }
            />
    }

    getEditOnePopup()
    {
        let popup = this.state.popups.editOnePopup;

        return <ModalEditArticle 
            open={popup.isOpen}
            onClose={ () => { this.changeEditPopupState(false) } } 
            onOk={ () => { this.onEdit() } }
            title={ popup.content.title }
            body={ popup.content.body }
            onChangeTitle={ (title) => { this.onChangeEditPopupTitle(title) } }
            onChangeBody={ (body)=>{ this.onChangeEditPopupBody(body) } }
            />
    }

    onChangeEditPopupTitle(title)
    {
        let popups = {...this.state.popups };
        popups.editOnePopup.content.title = title;
        this.setState({popups});
    }

    onChangeEditPopupBody(body)
    {
        let popups = {...this.state.popups };
        popups.editOnePopup.content.body = body;
        this.setState({popups});
    }

    async onOne(id)
    {
        let article = await ArticleModel.one(id);

        let popups = {...this.state.popups};
        let viewOnePopup = popups.viewOnePopup;
        viewOnePopup.isOpen = true;
        viewOnePopup.content = { ...article };

        this.setState( {popups} );
    }

    openViewModal(id)
    {
        let state = {...this.state};
        let articles = state.articles;
        let editOnePopupContent = state.popups.editOnePopup.content;
        let article = articles[id];

        editOnePopupContent.title = article.title;
        editOnePopupContent.body = article.body;
        
        state.popups.editOnePopup.content = editOnePopupContent;
        state.popups.editOnePopup.isOpen = true;
        state.popups.editOnePopup.article_id = id;
        
        this.setState(state);
    }

    render()
    {
        let articles = <ArticlesAll 
                        articles={this.state.articles} 
                        onDelete={ (id) =>{ this.onDelete(id); } } 
                        openEditModal={ (id)=>{ this.openViewModal(id) } }
                        /> 
        let addArticle = this.getAddOnePopup();
        let viewOnePopup = this.getViewOnePopup();
        let getEditOnePopup = this.getEditOnePopup();

        return <div>
            {viewOnePopup}
            {addArticle}
            {getEditOnePopup}
            <Typography variant="h3" component="h1">
                Товары - { this.state.articles.length }
                <Button onClick={ () => { this.changeAddPopupState(true); } } variant="outlined" color="primary" size="small">Добавить продукт</Button>
                <div>
                    <TextField
                        id="one_id"
                        variant="outlined"
                        label="Просмотр по ID"
                        onChange={ (e) => this.onOne( e.target.value ) }
                        />
                </div>
            </Typography>
            {articles}
        </div>
    }

}

export default Articles;