export {};
declare var BasePath: string

declare global {
    interface Window { BasePath: string; }
}
