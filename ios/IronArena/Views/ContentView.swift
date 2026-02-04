import SwiftUI

struct ContentView: View {
    @StateObject private var viewModel = MainViewModel()
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text("Welcome to Iron Arena")
                    .font(.largeTitle)
                    .fontWeight(.bold)
                
                Text(viewModel.welcomeMessage)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding()
                
                Button(action: {
                    viewModel.onGetStarted()
                }) {
                    Text("Get Started")
                        .font(.headline)
                        .foregroundColor(.white)
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color.blue)
                        .cornerRadius(10)
                }
                .padding(.horizontal)
            }
            .navigationTitle("Iron Arena")
        }
    }
}

#Preview {
    ContentView()
}
