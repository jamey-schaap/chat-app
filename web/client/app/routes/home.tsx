import type { Route } from "./+types/home";
import axios from "axios";
import type ChatMessage from "~/models/chat-message";
import { Chat } from "~/components/chat/chat.component";

export function meta({}: Route.MetaArgs) {
	return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export const clientLoader = async () =>
	axios.get<ChatMessage[]>("http://localhost:8080/chats").then((response) => response.data);

export default function Home({ loaderData }: Route.ComponentProps) {
	console.log(loaderData);
	return <Chat chatMessages={loaderData} onSendMessage={onSendMessage} />;
}

const onSendMessage = (message: string) => {
	console.log(message);
	axios.post("http://localhost:8080/chats", { message, id: "2", userId: "1" });
};
