import React from "react";
import Modal from "./modal-component"
import {Typography} from "@material-ui/core";

class ModalViewArticle extends Modal
{
    constructor( obj )
    {
        super(obj);
    }

    render()
    {

        return <div>
            <Modal 
                title="Просмотр товара"
                open={this.props.open}
                onClose={ this.props.onClose } 
                onOk={ this.props.onOk } >
                    <Typography  variant="h5" component="h2">{ this.props.title }</Typography>
                    <Typography  color="textSecondary">{ this.props.body }</Typography>
            </Modal>
        </div>
    }
}

export default ModalViewArticle;