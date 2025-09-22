import { create } from "zustand/react";
import type ChatMessage from "~/models/chat-message";

interface ChatStateSelectors {
	chatMessages: ChatMessage[];
}

interface ChatStateActions {
	setChatMessages: (chatMessages: ChatMessage[]) => void;
}

export const useChatStore = create<ChatStateSelectors & ChatStateActions>((set) => ({
	chatMessages: [],
	setChatMessages: (chatMessages: ChatMessage[]) => set(() => ({ chatMessages })),
}));
