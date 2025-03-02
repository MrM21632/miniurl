import "./App.css";

import { AppShell, MantineProvider, Title, rem } from "@mantine/core";
import "@mantine/core/styles.css";
import { NavBarSections } from "./components/NavBarSections";
import { useDisclosure, useHeadroom } from "@mantine/hooks";
import { Route, Routes } from "react-router-dom";
import { GeneratePage } from "./components/GeneratePage";
import { Link02Icon } from "@hugeicons-pro/core-stroke-standard";
import { HugeiconsIcon } from "@hugeicons/react";

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
            <HugeiconsIcon icon={Link02Icon} /> MiniURL
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
