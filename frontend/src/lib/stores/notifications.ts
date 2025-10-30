import { writable } from "svelte/store";

export type Notification = {
    id:  number,
    message: string,
    textColor: string,
    backgroundColor: string,
    onClick?: () => void,
    timeToSkip?: number,
    fading: boolean
}

const createPersistedNotifications = () => {
    const stored = localStorage.getItem("notifications");
    const initialNotifications = stored ? JSON.parse(stored): [];
    
    const {set, subscribe, update} = writable<Notification[]>(initialNotifications);

    const delay = (seconds: number) => new Promise(resolve => setTimeout(resolve, seconds*1000));

    subscribe((value) => {
        localStorage.setItem("notifications", JSON.stringify(value));
    });

    const addNotification = async({message, textColor, backgroundColor, onClick, timeToSkip}: Notification) => {
        const id = Date.now();
        update((n) => [...n, {id, message, textColor, backgroundColor, onClick, timeToSkip, fading: false}]);

        if (timeToSkip) {
            await delay(timeToSkip);

            update((current) => {
                const index = current.findIndex((c) => c.id == id);

                if (index == -1) return current;
                const updated = [...current];

                updated[index] = {...updated[index], fading: true};

                return current;
            })

            await delay(1);

            deleteNotification(id);
        }
    }



    const deleteNotification = (id: number) => {
        update((current) => current.filter(c => c.id != id));
    }

    return {addNotification, deleteNotification, subscribe, set};
} 

export const notifications = createPersistedNotifications();
