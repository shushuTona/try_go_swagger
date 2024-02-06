import './App.css'
import {Task, Api} from './gen/api'
import useSWR from 'swr'

const fetcher = ( url: string ) => fetch( url ).then( ( r ) => {
    const api = new Api()
    api.task.getTask()
  return r.json()
} )

function Profile() {
  const { data, error, isLoading } = useSWR<Task[]>( 'http://localhost:8085/api/task', fetcher )

  if ( error ) return <div>failed to load</div>
  if ( isLoading ) return <div>loading...</div>
  if ( data === undefined ) return <div>task is empty</div>

  return <ul>{data.map((item) => (
    <li key={item.taskId}>
        <div>
            <span>{item.title}</span>：<span>{item.desc}</span>
        </div>
    </li>
  ))}</ul>
}

function App() {
  return (
    <>
      <h1>try_swagger</h1>
      <div>
        <Profile />
      </div>
    </>
  )
}

export default App
