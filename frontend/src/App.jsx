import React from 'react'
import { Text, Box, Flex } from "@chakra-ui/react"
import { ProgressBar, ProgressLabel, ProgressRoot } from "@/components/ui/progress"

function App() {
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
            <Text fontSize="sm">アップデートしています。しばらくお待ち下さい...</Text>
            <ProgressRoot w="300px" value={null} mt="10px">
                <ProgressLabel mb="2">
                    Updating...
                </ProgressLabel>
                <ProgressBar />
            </ProgressRoot>
        </Box>
    )
}

export default App
