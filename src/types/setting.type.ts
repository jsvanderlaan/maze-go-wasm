export interface TextSourceInput {
    text: string;
    height: number;
    outline: boolean;
}

export interface ImageSourceInput {
    image: Uint8Array;
    height: number;
    threshold: number;
}
