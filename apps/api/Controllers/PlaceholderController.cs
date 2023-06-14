using Microsoft.AspNetCore.Mvc;

namespace api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class PlaceholderController : ControllerBase
    {
        private readonly ILogger<PlaceholderController> _logger;

        public PlaceholderController(ILogger<PlaceholderController> logger)
        {
            _logger = logger;
        }

        [HttpGet(Name = "GetPlaceholderData")]
        public string Get()
        {
            return "Hello World!";
        }
    }
}