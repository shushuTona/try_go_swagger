import { MouseEventHandler, ChangeEventHandler, useState, useCallback } from 'react'
import './App.css'
import TaskList from './components/TaskList'
import TaskForm from './components/TaskForm'
import { Task, ErrorResponse } from './gen/api'
import { api } from './api'
import useSWR from 'swr'

function App() {
  const [title, setTitle] = useState("")
  const [desc, setDesc] = useState( "" )

  const getTaskFetcher = useCallback( () => api.task.getTask().then( ( r ) => r.data ), [] )
  const { data, error, isLoading, mutate } = useSWR<Task[], ErrorResponse>( '/task', getTaskFetcher )

  const inputTitleHandler: ChangeEventHandler<HTMLInputElement> = useCallback( ( e ) => setTitle( e.target.value ), [] )
  const inputDescHandler: ChangeEventHandler<HTMLInputElement> = useCallback( ( e ) => setDesc( e.target.value ), [] )
  const addTaskHandler: MouseEventHandler<HTMLButtonElement> = useCallback( async () => {
    const body: Task = {
      title: title,
      desc: desc
    }
    await api.task.addTask( body ).then( ( r ) => { console.log( r ) } )
    await mutate( [body])
  }, [title, desc, mutate] )

  return (
    <>
      <h1>try_swagger</h1>
      <TaskList
        data={data}
        error={error}
        isLoading={isLoading}
      />
      <TaskForm
        title={title}
        desc={desc}
        inputTitleHandler={inputTitleHandler}
        inputDescHandler={inputDescHandler}
        addTaskHandler={addTaskHandler}
      />
    </>
  )
}

export default App
