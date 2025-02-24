package com.example.compose_view.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.example.compose_view.ui.page.HomePage

@Composable
fun RouterView(modifier: Modifier = Modifier) {
    Column(
        modifier=modifier,
    ) {
        HomePage()
    }
}

@Preview(showBackground = true)
@Composable
fun RouterViewPreview() {
    RouterView()
}