import "./App.css";

import { hugeiconsLicense } from "@hugeicons/react-pro";
import { MantineProvider } from "@mantine/core";
import "@mantine/core/styles.css";
import { NavBar } from "./components/NavBar";

console.log(import.meta.env.VITE_HUGEICONS_LICENSE_KEY);
hugeiconsLicense(import.meta.env.VITE_HUGEICONS_LICENSE_KEY || "");

function App() {
  return (
    <MantineProvider>
      <NavBar />
    </MantineProvider>
  );
}

export default App;
