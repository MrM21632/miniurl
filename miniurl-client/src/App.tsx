import "./App.css";

import { Link02Icon, hugeiconsLicense } from "@hugeicons/react-pro";
import { AppShell, MantineProvider, Title, rem } from "@mantine/core";
import "@mantine/core/styles.css";
import { NavBarSections } from "./components/NavBarSections";
import { useDisclosure, useHeadroom } from "@mantine/hooks";
import { Route, Routes } from "react-router-dom";
import { GeneratePage } from "./components/GeneratePage";

console.log(import.meta.env.VITE_HUGEICONS_LICENSE_KEY);
hugeiconsLicense(import.meta.env.VITE_HUGEICONS_LICENSE_KEY || "");

function App() {
  const [opened] = useDisclosure();
  const pinned = useHeadroom({ fixedAt: 120 });

  return (
    <MantineProvider>
      <AppShell
        header={{
          height: "3.5vh",
          collapsed: !pinned,
          offset: false,
        }}
        navbar={{
          width: 300,
          breakpoint: "sm",
          collapsed: { desktop: false, mobile: !opened },
        }}
      >
        <AppShell.Header>
          <Title order={1}>
            <Link02Icon color="var(--mantine-color-blue-4)" /> MiniURL
          </Title>
        </AppShell.Header>
        <AppShell.Navbar zIndex={0}>
          <NavBarSections />
        </AppShell.Navbar>
        <AppShell.Main pt={`calc(${rem(40)} + var(--mantine-spacing-md))`}>
          <Routes>
            <Route path="/" element={<GeneratePage />} />
          </Routes>
        </AppShell.Main>
      </AppShell>
    </MantineProvider>
  );
}

export default App;
