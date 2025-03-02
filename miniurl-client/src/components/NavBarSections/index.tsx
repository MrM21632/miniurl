import { PlusSignSquareIcon, Route01Icon } from "@hugeicons-pro/core-stroke-standard";
import { HugeiconsIcon } from "@hugeicons/react";
import { useState } from "react";

import { Stack } from "@mantine/core";
import "./index.css";

const data = [
  { link: "", label: "Generate", icon: <HugeiconsIcon icon={PlusSignSquareIcon} /> },
  { link: "", label: "Navigate", icon: <HugeiconsIcon icon={Route01Icon} /> },
];

export function NavBarSections() {
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
      {item.icon}
      <span>{item.label}</span>
    </a>
  ));

  return (
    <Stack
      align="stretch"
      justify="flex-start"
      gap="md"
      style={{ marginTop: "5vh", marginLeft: "2vh", marginRight: "1vh" }}
    >
      {links}
    </Stack>
  );
}
