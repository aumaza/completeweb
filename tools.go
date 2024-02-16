package main

import(
   "time"
)

func momentaryMessage(msg string) string {

    currentMessage := make(chan string, 1);

    text := msg;
    currentMessage <- text;

    time.Sleep(2 * time.Second);
    return text;
}
