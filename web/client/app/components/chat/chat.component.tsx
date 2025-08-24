import logoLight from "~/components/chat/logo-light.svg";
import logoDark from "~/components/chat/logo-dark.svg";
import SendIconSvg from "~/icons/send.icon.svg";
import type ChatMessage from "~/models/chat-message";
import { useState } from "react";

type Props = { chatMessages: ChatMessage[]; onSendMessage: (message: string) => Promise<void> | void };

export const Chat = ({ chatMessages, onSendMessage }: Props) => {
	const [message, setMessage] = useState<string>("");

	const sendMessage = async () => {
		if (message.trim() === "") {
			return;
		}

		await onSendMessage(message);
		setMessage("");
	};

	return (
		<main className="flex items-center justify-center pt-16 pb-4">
			<div className="flex-1 flex flex-col items-center gap-16 min-h-0">
				<header className="flex flex-col items-center gap-9">
					<div className="w-[500px] max-w-[100vw] p-4">
						<img src={logoLight} alt="React Router" className="block w-full dark:hidden" />
						<img src={logoDark} alt="React Router" className="hidden w-full dark:block" />
					</div>
				</header>
				<div className="max-w-[400px] w-full space-y-6 px-4">
					<div className="rounded-3xl border border-gray-200 p-6 dark:border-gray-700 space-y-4">
						{chatMessages.map((chatMessage) => (
							<p key={chatMessage.id} className="leading-6 text-gray-700 dark:text-gray-200 text-center">
								{chatMessage.userId}: {chatMessage.message}
							</p>
						))}
					</div>
					<div className="flex items-center justify-center gap-4">
						<input
							type="text"
							placeholder="Type your message..."
							className="w-full rounded-3xl border border-gray-200 p-4 dark:border-gray-700"
							value={message}
							onChange={(event) => {
								setMessage(event.target.value);
							}}
							onKeyDown={async (event) => {
								if (event.key === "Enter") {
									await sendMessage();
								}
							}}
						/>
						<button
							type="button"
							className="cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-800 rounded-full p-4"
							onClick={() => sendMessage()}
						>
							<SendIconSvg />
						</button>
					</div>
				</div>
			</div>
		</main>
	);
};
