import './App.css'
import useSWR from 'swr'

const fetcher = ( url: string ) => fetch( url ).then( ( r ) => {
  console.log("r", r);

  return r.json()
} )

function Profile() {
  const { data, error, isLoading } = useSWR( 'http://localhost:8081/api/task', fetcher )

  console.log(error);

  if ( error ) return <div>failed to load</div>
  if ( isLoading ) return <div>loading...</div>
  return <div>hello {data}!</div>
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
