import { MouseEventHandler, ChangeEventHandler } from 'react'

type TaskFormProps = {
    title: string
    desc: string
    inputTitleHandler: ChangeEventHandler<HTMLInputElement>
    inputDescHandler: ChangeEventHandler<HTMLInputElement>
    addTaskHandler: MouseEventHandler<HTMLButtonElement>
}

const TaskForm = ( props: TaskFormProps ) => {
    return (
        <form>
            <p><label>title : <input type="text" name="title" value={props.title} onChange={props.inputTitleHandler} /></label></p>
            <p><label>desc : <input type="text" name="desc" value={props.desc} onChange={props.inputDescHandler} /></label></p>
            <button type="button" onClick={props.addTaskHandler}>add task</button>
        </form>
    )
}

export default TaskForm
