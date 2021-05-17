import React from "react";
import Modal from "./modal-component"
import {TextField} from "@material-ui/core";

class ModalAddArticle extends Modal
{
    constructor( obj )
    {
        super(obj);

        this.state = {
            model: this.props.title,
            company: this.props.content
        }
    }

    onChangeTitle(title)
    {
        let state = {...this.state};
        state.model = title;
        this.setState(state);
    }

    onChangeContent(content)
    {
        let state = {...this.state};
        state.company = content;
        this.setState(state);
    }


    render()
    {
        return <Modal 
            title="Добавить продукт"
            open={this.props.open}
            onClose={ this.props.onClose }
            onOk={ () => { this.props.onOk(this.state) } } >
                <form>
                    <TextField
                        id="title"
                        variant="outlined"
                        label="Модель"
                        onChange={ (e) => { this.onChangeTitle(e.target.value) } }
                        />
                    <hr/>
                    <TextField
                        id="content"
                        variant="outlined"
                        label="Компания"
                        multiline
                        rowsMax={4}
                        onChange={ (e) => { this.onChangeContent(e.target.value) } }
                        />
                </form>
            </Modal>
    }
}

export default ModalAddArticle;