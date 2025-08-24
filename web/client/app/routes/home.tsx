import type { Route } from "./+types/home";
import axios from "axios";
import type ChatMessage from "~/models/chat-message";
import { Chat } from "~/components/chat/chat.component";
import { useRef } from "react";

export function meta({}: Route.MetaArgs) {
	return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export const clientLoader = async () =>
	axios.get<ChatMessage[]>("http://localhost:8080/chats").then((response) => response.data);

export default function Home({ loaderData }: Route.ComponentProps) {
	// const wsRef = useRef<WebSocket | null>(null);
	const socket = new WebSocket("ws://localhost:8080/echo");

	socket.onopen = () => {
		console.log("Connected to the WebSocket server");
	};
	socket.onclose = () => {
		console.log("Disconnected from the WebSocket server");
	};
	socket.onmessage = (event: MessageEvent) => {
		console.log("Received from WebSocket server: " + event.data);
	};

	const onSendMessage = async (message: string) => {
		console.log("Send from client: " + message);
		// await axios.post("http://localhost:8080/chats", { message, id: "2", userId: "1" });
		socket.send(message);
	};

	return <Chat chatMessages={loaderData} onSendMessage={onSendMessage} />;
}
