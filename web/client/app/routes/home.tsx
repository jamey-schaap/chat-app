import type { Route } from "./+types/home";
import axios from "axios";
import type ChatMessage from "~/models/chat-message";
import { Chat } from "~/components/chat/chat.component";
import { useChatStore } from "~/hooks/stores/useChatStore";
import { useEffect } from "react";

export function meta({}: Route.MetaArgs) {
	return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export const loader = async () =>
	axios.get<ChatMessage[]>("http://localhost:8080/chats").then((response) => response.data);

export default function Home({ loaderData: initialChatMessages }: Route.ComponentProps) {
	const chatMessages = useChatStore((state) => state.chatMessages);
	const setChatMessages = useChatStore((state) => state.setChatMessages);

	useEffect(() => {
		setChatMessages(initialChatMessages);
	}, [initialChatMessages]);

	const onSendMessage = (message: string) => axios.post("http://localhost:8080/chats", JSON.stringify({ message }));
	return <Chat chatMessages={initialChatMessages} onSendMessage={onSendMessage} />;
}
