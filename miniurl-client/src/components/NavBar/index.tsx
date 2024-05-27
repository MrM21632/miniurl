import {
  Link02Icon,
  PlusSignSquareIcon,
  Route01Icon,
} from "@hugeicons/react-pro";
import { useState } from "react";

import { Group, Title } from "@mantine/core";
import "./index.css";

const data = [
  { link: "", label: "Generate", icon: PlusSignSquareIcon },
  { link: "", label: "Navigate", icon: Route01Icon },
];

export function NavBar() {
  const [active, setActive] = useState("Generate");

  const links = data.map((item) => (
    <a
      className="link"
      data-active={item.label === active || undefined}
      href={item.link}
      key={item.label}
      onClick={(event) => {
        event.preventDefault();
        setActive(item.label);
      }}
    >
      <item.icon className="link-icon" strokeWidth={1.5} />
      <span>{item.label}</span>
    </a>
  ));

  return (
    <nav className="navbar">
      <div className="navbar-main">
        <Group className="header" justify="space-between">
          <Link02Icon color="var(--mantine-color-blue-5)" strokeWidth={2.5} />
          <Title c="blue.8" order={1}>
            MiniURL
          </Title>
        </Group>
        {links}
      </div>
    </nav>
  );
}
