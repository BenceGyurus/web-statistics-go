export default interface SelectorParams {
    options : string[];
    selected : string;
    onSelect : (value: string) => void;
}