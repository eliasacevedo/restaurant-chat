export interface Message {
    id?: number,
    text: string,
    user: string,
    time: Date;
    classification?: string[];
}