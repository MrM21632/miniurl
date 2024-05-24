import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import '@mantine/core/styles.css'
import { MantineProvider } from '@mantine/core'
import { Upload01Icon, hugeiconsLicense } from '@hugeicons/react-pro'

console.log(import.meta.env.VITE_HUGEICONS_LICENSE_KEY)
hugeiconsLicense(import.meta.env.VITE_HUGEICONS_LICENSE_KEY || '')

function App() {
  const [count, setCount] = useState(0)

  return (
    <MantineProvider>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          <Upload01Icon /> count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </MantineProvider>
  )
}

export default App
