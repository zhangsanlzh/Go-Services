using System;
using RestSharp;
using System.IO;

namespace client
{
    class Program
    {
        static void Main(string[] args)
        {
            string fileName="Microsoft_.NET_logo.png";

            var client = new RestClient("http://localhost:8000/api/download/");
            IRestRequest request = new RestRequest(fileName, Method.GET);

            var bytes = client.DownloadData(request);
            var response = client.Execute(request);
            File.WriteAllBytes(fileName, bytes);

            Console.WriteLine(response.ResponseStatus);            
        }
    }
}
