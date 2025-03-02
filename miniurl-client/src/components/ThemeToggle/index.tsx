import { Moon02Icon, Sun03Icon } from "@hugeicons-pro/core-stroke-standard";
import { HugeiconsIcon } from "@hugeicons/react";

function ThemeToggle() {
  return (
    <label className="flex cursor-pointer gap-2">
      <HugeiconsIcon icon={Sun03Icon} />
      <input
        type="checkbox"
        value="darknord"
        className="toggle theme-controller"
      />
      <HugeiconsIcon icon={Moon02Icon} />
    </label>
  );
}

export default ThemeToggle;
