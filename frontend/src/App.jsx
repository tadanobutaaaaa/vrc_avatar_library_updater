import React, { useState } from 'react'
import { EventsOn } from "../wailsjs/runtime/runtime"
import { Text, Box, Alert } from "@chakra-ui/react"
import { ProgressBar, ProgressLabel, ProgressRoot } from "@/components/ui/progress"

function App() {
    const [isEventOn, setIsEventOn] = useState(false)
    EventsOn("error", () => {
        setIsEventOn(true)
    })

    return (
        <Box 
            height="100vh"      // ビューポート全体の高さを設定
            width="100vw"       // ビューポート全体の幅を設定
            overflow="hidden"   // スクロールバーを非表示にする
            display="flex" 
            flexDirection="column"
            alignItems="center"
            justifyContent="center"
        >
            {isEventOn ? (
                <Alert.Root status="error">
                    <Alert.Indicator />
                    <Alert.Content>
                        <Alert.Title>予期せぬエラーが発生しました</Alert.Title>
                        <Alert.Description>
                            「VRC-Avatar-Library」が起動していないか確認し、再度お試しください。
                        </Alert.Description>
                    </Alert.Content>
                </Alert.Root>
            ) : (
                <>
                    <Text fontSize="sm">アップデートしています。しばらくお待ち下さい...</Text>
                    <ProgressRoot w="300px" value={null} mt="10px">
                        <ProgressLabel mb="2">
                            Updating...
                        </ProgressLabel>
                        <ProgressBar />
                    </ProgressRoot>
                </>
            )}
        </Box>
    )
}

export default App
