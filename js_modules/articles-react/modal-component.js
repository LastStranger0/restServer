import React from "react";
import Article from "./article-component";
import {Button, Dialog, DialogContent, DialogTitle, DialogActions} from "@material-ui/core";

class Modal extends React.Component
{
    constructor( obj )
    {
        super(obj);
    }

    render()
    {
        return <div>
            <Dialog
                open={this.props.open}
                onClose={ this.props.onClose }
            >
            <DialogTitle>{this.props.title}</DialogTitle>
            <DialogContent>{ this.props.children }</DialogContent>
            <DialogActions>
                <Button onClick={ this.props.onOk } color="primary">Ок</Button>
                <Button onClick={ this.props.onClose } color="secondary">Закрыть</Button>
            </DialogActions>
            </Dialog>
        </div>
    }
}

export default Modal;