import { createContext, type Dispatch, type ReactNode, useContext, useReducer } from "react";
import type ChatMessage from "~/models/chat-message";

const initialChatMessages: ChatMessage[] = [];
const ChatMessagesContext = createContext<ChatMessage[]>(initialChatMessages);
const ChatMessagesDispatchContext = createContext<Dispatch<ChatMessageAction>>(() => {
	throw new Error("ChatMessagesDispatchContext must be used withing ChatMessagesProvider");
});

export function ChatMessagesProvider({ children }: { children: ReactNode }) {
	const [chatMessages, dispatch] = useReducer(chatMessagesReducer, initialChatMessages);

	return (
		<ChatMessagesContext value={chatMessages}>
			<ChatMessagesDispatchContext value={dispatch}>{children}</ChatMessagesDispatchContext>
		</ChatMessagesContext>
	);
}

export function useChatMessages() {
	const context = useContext(ChatMessagesContext);
	if (!context) {
		throw new Error("useChatMessages must be used withing ChatMessagesProvider");
	}

	return context;
}

export function useChatMessagesDispatch() {
	return useContext(ChatMessagesDispatchContext);
}

const ChatMessageActionType = {
	ChatMessageCreated: 0,
} as const;

type ChatMessageActionType = (typeof ChatMessageActionType)[keyof typeof ChatMessageActionType];

type ChatMessageAction = {
	type: typeof ChatMessageActionType.ChatMessageCreated;
	message: string;
};

function chatMessagesReducer(state: ChatMessage[], action: ChatMessageAction) {
	switch (action.type) {
		case ChatMessageActionType.ChatMessageCreated: {
			// create
			return [...state];
		}
		default: {
			throw Error("Unknown action: " + action.type);
		}
	}
}
