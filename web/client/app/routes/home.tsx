import type { Route } from "./+types/home";
import axios from "axios";
import type ChatMessage from "~/models/chat-message";
import { Welcome } from "~/welcome/welcome";

export function meta({}: Route.MetaArgs) {
	return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export const clientLoader = async () =>
	axios.get<ChatMessage[]>("http://localhost:8080/chats").then((response) => response.data);

export default function Home({ loaderData }: Route.ComponentProps) {
	console.log(loaderData);
	return <Welcome />;
}
