export type Notification = {
    id:  number,
    message: string,
    textColor: string,
    backgroundColor: string,
    onClick?: () => void,
    timeToSkip?: number,
    fading: boolean
}