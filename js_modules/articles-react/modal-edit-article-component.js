import React from "react";
import Modal from "./modal-component"
import {TextField} from "@material-ui/core";

class ModalEditArticle extends Modal
{
    constructor( obj )
    {
        super(obj);
    }

    render()
    {
        return <Modal 
            title="Редактировать статью"
            open={this.props.open}
            onClose={ this.props.onClose }
            onOk={ this.props.onOk } >
                <form>
                    <TextField
                        variant="outlined"
                        label="Заголовок"
                        onChange={ (e) => { this.props.onChangeTitle(e.target.value) } }
                        value={this.props.title}
                        />
                    <hr/>
                    <TextField
                        variant="outlined"
                        label="Контент"
                        multiline
                        rowsMax={4}
                        onChange={ (e) => { this.props.onChangeBody(e.target.value) } }
                        value={this.props.body}
                        />
                </form>
            </Modal>
    }
}

export default ModalEditArticle;