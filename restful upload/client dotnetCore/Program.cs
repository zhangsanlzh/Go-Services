using System;
using RestSharp;
using System.IO;

namespace client
{
    class Program
    {
        static void Main(string[] args)
        {             
            UploadFile("files/logo.png");
        }

        public static void UploadFile(string filePath)
        {
            RestClient client = new RestClient("http://localhost:8000");
            IRestRequest request = new RestRequest("/upload", Method.POST);

            request.AddFile("file", filePath);
            IRestResponse response = client.Execute(request);
            Console.WriteLine(response.Content);
        } 
    }    
}
