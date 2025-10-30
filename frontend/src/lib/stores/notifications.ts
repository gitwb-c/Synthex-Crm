import { browser } from "$app/environment";
import type { Notification } from "$lib/models/types/notification";
import { writable } from "svelte/store";


const createPersistedNotifications = () => {
    const stored = browser? localStorage.getItem("notifications"):null;
    const initialNotifications = stored ? JSON.parse(stored): [];
    
    const {set, subscribe, update} = writable<Notification[]>(initialNotifications);

    const delay = (seconds: number) => new Promise(resolve => setTimeout(resolve, seconds*1000));

    subscribe((value) => {
        if (browser) localStorage.setItem("notifications", JSON.stringify(value));
    });

    const addNotification = async({message, textColor, backgroundColor, onClick, timeToSkip}: Omit<Notification, "id" | "fading">) => {
        const id = Date.now();
        update((n) => [...n, {id, message, textColor, backgroundColor, onClick, timeToSkip, fading: false}]);

        if (timeToSkip) {
            await delay(timeToSkip);

            update((current) => {
                const index = current.findIndex((c) => c.id == id);

                if (index == -1) return current;
                const updated = [...current];

                updated[index] = {...updated[index], fading: true};

                return updated;
            })

            await delay(1);

            deleteNotification(id);
        }
    }



    const deleteNotification = async(id: number) => {
        update((current) => {
                const index = current.findIndex((c) => c.id == id);

                if (index == -1) return current;
                const updated = [...current];

                updated[index] = {...updated[index], fading: true};

                return updated;
            });
        await delay(1);
        update((current) => current.filter(c => c.id != id));
    }

    return {addNotification, deleteNotification, subscribe, set};
} 

export const notifications = createPersistedNotifications();
