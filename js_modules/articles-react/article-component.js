import React from "react";
// eslint-disable-next-line no-unused-vars
import {Card, Typography, CardActions, Button, CardContent} from '@material-ui/core';

class Article extends React.Component
{
    constructor(obj)
    {
        super(obj);

        this.onDelete = "onDelete" in this.props ? this.props.onDelete : function(){};
    }

    render()
    {
        return <Card variant="outlined">
            <CardContent>
                <Typography  variant="h5" component="h2">{this.props.title}</Typography>
                <Typography  color="textSecondary">{this.props.body}</Typography>
            </CardContent>
            <CardActions>
                <Button onClick={ this.onDelete } size="small">Удалить</Button>
                <Button onClick={ this.props.openEditModal } size="small">Редактировать</Button>
            </CardActions>
        </Card>
    }
}

export default Article;