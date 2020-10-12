import { getEmbeddedVariable } from "./Configuration";

export function useTheme() {
    return getEmbeddedVariable("theme-name");
}

export function usePrimaryColor() {
    return getEmbeddedVariable("theme-primarycolor");
}

export function useSecondaryColor() {
    return getEmbeddedVariable("theme-secondarycolor");
}
