import { Task, ErrorResponse } from '../gen/api'

type TaskListProps = {
    data: Task[] | undefined
    error: ErrorResponse | undefined
    isLoading: boolean
}

const TaskList = ( props: TaskListProps ) => {
    if ( props.error || props.data === undefined ) return <div>failed to load</div>
    if ( props.isLoading ) return <div>loading...</div>
    if ( props.data.length === 0 ) return <div>task is empty</div>

    const tasks = props.data.map( ( item ) => {
        if ( item.taskId ) {
            return (
                <li key={item.taskId}>
                    <div>
                        <span>{item.title}</span>：<span>{item.desc}</span>
                    </div>
                </li>
            )
        }
    })

    return <ul>{tasks}</ul>
}

export default TaskList
